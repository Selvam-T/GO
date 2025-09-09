FROM golang:1.25

WORKDIR /usr/src/app

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        nano \
        curl \
        git \
        postgresql \
        postgresql-contrib \
        && rm -rf /var/lib/apt/lists/*
        
# Create directory for SQL scripts
RUN mkdir -p /docker-entrypoint-initdb.d

# Copy init SQL and entrypoint
COPY init.sql /docker-entrypoint-initdb.d/init.sql
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Expose ports
EXPOSE 8080 5432

# Run entrypoint
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

#CMD ["bash"]
