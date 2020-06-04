FROM gcr.io/kubeflow-images-public/tensorflow-2.1.0-notebook-gpu:1.0.0
USER root
RUN pip3 install \
  # google-cloud-storage==1.25.0 \
  tensorflow-datasets==2.1.0 \
  kubeflow-metadata==0.3.1 \
  pandas==1.0.3
ENV NB_PREFIX /
