#!/bin/bash -e

set -o pipefail

MIRROR_MSG="Mirrored from envoyproxy/envoy"
SRCS=(envoy contrib)
GO_TARGETS=(@envoy_api//...)
IMPORT_BASE="github.com/envoyproxy/go-control-plane"
COMMITTER_NAME="update-envoy[bot]"
COMMITTER_EMAIL="135279899+update-envoy[bot]@users.noreply.github.com"
ENVOY_SRC_DIR="${ENVOY_SRC_DIR:-}"


if [[ -z "$ENVOY_SRC_DIR" ]]; then
    echo "ENVOY_SRC_DIR not set, it should point to a cloned Envoy repo" >&2
    exit 1
elif [[ ! -d "$ENVOY_SRC_DIR" ]]; then
    echo "ENVOY_SRC_DIR ($ENVOY_SRC_DIR) not found, did you clone it?" >&2
    exit 1
fi


build_protos () {
    echo "Building go protos ..."
    cd "${ENVOY_SRC_DIR}" || exit 1
    ./ci/do_ci.sh api.go
    cd - || exit 1
}

get_last_envoy_sha () {
    git log \
        --grep="$MIRROR_MSG" -n 1 \
        | grep "$MIRROR_MSG" \
        | tail -n 1 \
        | sed -e "s#.*$MIRROR_MSG @ ##"
}

sync_protos () {
    local src envoy_src
    echo "Syncing go protos ..."
    for src in "${SRCS[@]}"; do
        envoy_src="${ENVOY_SRC_DIR}/build_go/${src}"
        rm -rf "$src"
        echo "Copying ${envoy_src} -> ${src}"
        cp -a "$envoy_src" "$src"
        git add "$src"
    done
    make tidy-all
    git add $(find -type f -name 'go.sum' -o -name 'go.mod')
}

commit_changes () {
    local last_envoy_sha changes changed
    echo "Committing changes ..."
    changed="$(git diff HEAD --name-only | grep -v envoy/COMMIT || :)"
    if [[ -z "$changed" ]]; then
        echo "Nothing changed, not committing"
        return
    fi
    last_envoy_sha="$(get_last_envoy_sha)"
    echo "Latest Envoy SHA: ${last_envoy_sha}"
    changes="$(git -C "${ENVOY_SRC_DIR}" rev-list "${last_envoy_sha}"..HEAD)"
    echo "Changes detected: "
    echo "$changes"
    latest_commit="$(git -C "${ENVOY_SRC_DIR}" rev-list HEAD -n1)"
    echo "Latest commit: ${latest_commit}"
    echo "$latest_commit" > envoy/COMMIT
    git config user.email "$COMMITTER_EMAIL"
    git config user.name "$COMMITTER_NAME"
    git add envoy contrib
    git commit --allow-empty -s -m "${MIRROR_MSG} @ ${latest_commit}"
    git push origin main
}


build_protos
sync_protos
commit_changes
