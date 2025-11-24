#!/bin/sh
java ${JavaOptions} ${APM_OPTIONS} ${JVM_OPTIONS} -Dxxl.job.executor.ip=${SERVER_IP} -jar /opt/${PackageName}