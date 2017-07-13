# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: application.proto

require 'google/protobuf'

require 'soapbox_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "soapbox.Application" do
    optional :id, :int32, 1
    optional :name, :string, 2
    optional :description, :string, 3
    optional :external_dns, :string, 4
    optional :github_repo_url, :string, 5
    optional :dockerfile_path, :string, 6
    optional :entrypoint_override, :string, 7
    optional :type, :enum, 8, "soapbox.ApplicationType"
    optional :created_at, :string, 9
    optional :slug, :string, 10
    optional :internal_dns, :string, 11
  end
  add_message "soapbox.ListApplicationResponse" do
    repeated :applications, :message, 1, "soapbox.Application"
  end
  add_message "soapbox.GetApplicationRequest" do
    optional :id, :int32, 1
  end
  add_enum "soapbox.ApplicationType" do
    value :SERVER, 0
    value :CRONJOB, 1
  end
end

module Soapbox
  Application = Google::Protobuf::DescriptorPool.generated_pool.lookup("soapbox.Application").msgclass
  ListApplicationResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("soapbox.ListApplicationResponse").msgclass
  GetApplicationRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("soapbox.GetApplicationRequest").msgclass
  ApplicationType = Google::Protobuf::DescriptorPool.generated_pool.lookup("soapbox.ApplicationType").enummodule
end