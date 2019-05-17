provider "divvycloud" {
  address = "https://yourdivvyhost/v2/"
  username = "username"
  password = "password"
}

resource "divvycloud_organization" "organization" {
  name = "myorg"
}

resource "divvycloud_aws_cloud_account_sts" "consumer" {
  name = "consumer-account"
  account_id = "123456"
  role_arn = "arn:aws:iam::123456:role/divvycloud"
  organization_id = "${divvycloud_event_driven_harvesting.event_driven_harvesting.organization_id}"
}

resource "divvycloud_aws_cloud_account_sts" "producer" {
  name = "producer-account"
  account_id = "654321"
  role_arn = "arn:aws:iam::654321:role/divvycloud"
  organization_id = "${divvycloud_event_driven_harvesting.event_driven_harvesting.organization_id}"
}

resource "divvycloud_event_driven_harvesting" "event_driven_harvesting" {
  organization_id = "${divvycloud_organization.organization.id}"
  enabled = true
}

resource "divvycloud_event_driven_harvesting_consumer" "consumer" {
  cloud_id = "${divvycloud_aws_cloud_account_sts.consumer.cloud_id}"
  organization_id = "${divvycloud_event_driven_harvesting.event_driven_harvesting.organization_id}"
}

resource "divvycloud_event_driven_harvesting_producer" "producer" {
  cloud_id = "${divvycloud_aws_cloud_account_sts.producer.cloud_id}"
  consumer_cloud_id = "${divvycloud_event_driven_harvesting_consumer.consumer.cloud_id}"
  organization_id = "${divvycloud_event_driven_harvesting.event_driven_harvesting.organization_id}"
}
