data "aws_iam_policy_document" "dyplomat-policy" {
  statement {
    actions = [
      "sts:AssumeRoleWithWebIdentity"
    ]

    principals {
      type = "Federated"
      identifiers = [vars.oidc_provider_demo2]
    }
    
    effect = "Allow"
  }

  statement {
    actions = [
      "sts:AssumeRoleWithWebIdentity"
    ]

    principals {
      type = "Federated"
      identifiers = [vars.oidc_provider_demo1]
    }
    
    effect = "Allow"
  }
}

resource "aws_iam_role" "dyploymat" {
  name = "dyplomat"

  assume_role_policy = "${data.aws_iam_policy_document.dyplomat-policy.json}"

  tags = {
    service = "dyplomat"
  }
}
