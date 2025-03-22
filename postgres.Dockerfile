FROM postgres:latest

# Install build dependencies and pgvector
RUN apt-get update \
    && apt-get install -y \
        postgresql-server-dev-all \
        git \
        build-essential \
    && git clone --branch v0.6.0 https://github.com/pgvector/pgvector.git \
    && cd pgvector \
    && make \
    && make install \
    && cd .. \
    && rm -rf pgvector \
    && apt-get remove -y postgresql-server-dev-all git build-essential \
    && apt-get autoremove -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Create a script to enable the extension and make it executable
RUN echo "CREATE EXTENSION IF NOT EXISTS vector;" > /docker-entrypoint-initdb.d/10-enable-vector.sql