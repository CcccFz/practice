#!/bin/bash
ver="5.7"
tag="5.7.20-1.el7"

#wget https://cdn.mysql.com//Downloads/MySQL-$ver/mysql-$tag.x86_64.rpm-bundle.tar
#wget ftp://195.220.108.108/linux/centos/7.4.1708/os/x86_64/Packages/libaio-0.3.109-13.el7.x86_64.rpm

#tar -xvf mysql-$tag.x86_64.rpm-bundle.tar
rpm -ivh libaio-0.3.109-13.el7.x86_64.rpm
rpm -ivh mysql-community-common-$tag.x86_64.rpm
rpm -ivh mysql-community-libs-$tag.x86_64.rpm
rpm -ivh mysql-community-client-$tag.x86_64.rpm
rpm -ivh mysql-community-server-$tag.x86_64.rpm
rpm -ivh mysql-community-devel-$tag.x86_64.rpm
rpm -ivh mysql-community-libs-compat-$tag.x86_64.rpm

