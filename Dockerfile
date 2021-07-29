# The default obtained 6/4 from /universe/docker-images/ubuntu/18.04/BUILD
ARG IMAGE_SHA=sha256:6058c5e6bd1c60a065f5197d705386c84c62c4ab806ef51a7342a3237fcd1e5d
# Set base image to internal ubuntu image with image sha version generated in universe/thirdparty/showfast/build_container.sh
FROM registry.dev.databricks.com/universe/db-ubuntu-18.04@${IMAGE_SHA} AS final

# Install Necessary Tools
RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get autoremove -y \
    && apt-get clean \
    && rm -rf /tmp/* /var/tmp/* \
    && rm -rf /var/lib/apt/lists/*

# Copy Binary
COPY bin/example /databricks/xds-go-cp
RUN chmod +x /databricks/xds-go-cp

# Port
EXPOSE 18000

# Run showfast
CMD ["/bin/bash", "-c", "/databricks/xds-go-cp -debug $@"]
