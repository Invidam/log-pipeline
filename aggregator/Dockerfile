FROM fluent/fluentd:v1.17-debian-1

# Use root user to install additional plugins
USER root

# Install the Kafka plugin for Fluentd
RUN apt-get update && \
    apt-get install -y gcc make && \
    gem install fluent-plugin-kafka --no-document && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Switch back to fluent user
USER fluent
