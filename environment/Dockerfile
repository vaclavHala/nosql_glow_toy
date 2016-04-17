FROM debian:latest
RUN apt-get -y update 

# MongoDB
RUN apt-get -y install mongodb
RUN mkdir -p /mongo/db
ADD products.json /mongo/products.json

EXPOSE 27017 27018 27019 28017

# Hadoop
RUN apt-get -y install default-jdk
ENV JAVA_HOME=/usr/bin/java

RUN apt-get -y install ssh
RUN ssh-keygen -t rsa -P "" -f /root/.ssh/id_rsa
RUN cat $HOME/.ssh/id_rsa.pub >> $HOME/.ssh/authorized_keys
# Avoid "trust host?" from ssh in start script
RUN echo "StrictHostKeyChecking no" >> /etc/ssh/ssh_config
RUN apt-get -y install rsync

EXPOSE 22

# for downloading hadoop itself
RUN apt-get -y install curl

RUN mkdir -p /hadoop/hdfs
RUN curl -o /hadoop/hadoop.tar.gz http://mirrors.sonic.net/apache/hadoop/common/hadoop-2.6.0/hadoop-2.6.0.tar.gz
RUN cd /hadoop && tar xvzf hadoop.tar.gz && cd /

ENV HADOOP_INSTALL=/hadoop/hadoop-2.6.0
ENV PATH=$PATH:$HADOOP_INSTALL/bin
ENV PATH=$PATH:$HADOOP_INSTALL/sbin
ENV HADOOP_MAPRED_HOME=$HADOOP_INSTALL
ENV HADOOP_COMMON_HOME=$HADOOP_INSTALL
ENV HADOOP_HDFS_HOME=$HADOOP_INSTALL
ENV YARN_HOME=$HADOOP_INSTALL
ENV HADOOP_COMMON_LIB_NATIVE_DIR=$HADOOP_INSTALL/lib/native
ENV HADOOP_OPTS="-Djava.library.path=$HADOOP_INSTALL/lib"
# real JAVA_HOME must be /usr/bin/java
# stupid hadoop puts the /bin/java at the end itself... 
RUN echo "export JAVA_HOME=/usr" >> $HADOOP_INSTALL/etc/hadoop/hadoop-env.sh
ADD core-site.xml $HADOOP_INSTALL/etc/hadoop/core-site.xml
ADD hdfs-site.xml $HADOOP_INSTALL/etc/hadoop/hdfs-site.xml
ADD hdfs-site.xml $HADOOP_INSTALL/etc/hadoop/mapred-site.xml

RUN mkdir -p /hadoop/hdfs/namenode
RUN mkdir -p /hadoop/hdfs/datanode

ADD products.csv /hadoop/products.csv

EXPOSE 12300 12311

# start script fires up servers, inits data etc.
ADD start /start
RUN chmod 700 /start
CMD bash -C '/start';'bash'