import docker
client = docker.from_env()



print client.containers.list()

print client.images.list()

client.images.pull("arifch2009/ubuntu-arif")

print client.images.list()
