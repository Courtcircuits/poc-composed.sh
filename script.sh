#!/bin/bash

# kctl commands to manager rights
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default
