# Based on ubuntu
FROM ubuntu:18.04
LABEL maintainers="Tuan L"
LABEL description="KubeEdge Trafficlight App"

# Copy from build directory
COPY trafficlight-app /trafficlight-app

# Define default command
ENTRYPOINT ["/trafficlight-app"]

# Run the executable
CMD ["trafficlight-app"]