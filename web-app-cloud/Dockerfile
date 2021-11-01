FROM ubuntu:18.04
LABEL maintainer="Tuan L"
LABEL description="Kube edge traffic light simu"

COPY kubeedge-lighttraffic-controller /kubeedge-lighttraffic-controller
COPY static /static
COPY views /views
COPY images /images

# Define default command
ENTRYPOINT ["/kubeedge-lighttraffic-controller"]

# Run executable
CMD ["kubeedge-lighttraffic-controller"]