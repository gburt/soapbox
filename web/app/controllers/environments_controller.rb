require 'environment_pb'

class EnvironmentsController < ApplicationController
  before_action :set_application, only: [:index, :create, :new, :show]

  def index
    app_id = params[:application_id].to_i
    req = Soapbox::ListEnvironmentRequest.new(application_id: app_id)
    res = $api_environment_client.list_environments(req)
    if res.environments.count == 0
      redirect_to new_application_environment_path
    else
      @environments = []
      res.environments.each do |env|
        begin
          latest_deploy = get_latest_deploy(app_id, env.id)
        rescue GRPC::NotFound
          latest_deploy = nil
        end
        @environments << [env, latest_deploy]
      end
    end
  end

  def new
    @form = CreateEnvironmentForm.new
  end

  def create
    @form = CreateEnvironmentForm.new(params[:environment])
    if @form.valid?
      env = Soapbox::Environment.new(application_id: params[:application_id].to_i, name: @form.name)
      $api_environment_client.create_environment(env)
      redirect_to application_environments_path
    else
      render :new
    end
  end

  def show
    env_id = params[:id].to_i
    @environment = get_environment(env_id)
  end

  def destroy
    req = Soapbox::DestroyEnvironmentRequest.new(id: params[:id].to_i)
    $api_environment_client.destroy_environment(req)
    redirect_to application_environments_path
  end

  private

  def set_application
    req = Soapbox::GetApplicationRequest.new(id: params[:application_id].to_i)
    @app = $api_client.get_application(req)
  end

  def get_environment(id)
    req = Soapbox::GetEnvironmentRequest.new(id: id)
    $api_environment_client.get_environment(req)
  end

  def get_latest_deploy(app_id, env_id)
    req = Soapbox::GetLatestDeploymentRequest.new(application_id: app_id, environment_id: env_id)
    $api_deployment_client.get_latest_deployment(req)
  end
end
