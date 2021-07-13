# fuzzton

## Prerequisites

Setup a Kubernetes cluster 

Install [Tekton Pipelines](https://github.com/tektoncd/pipeline)

```shell
kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
```

Install [Tekton CLI](https://github.com/tektoncd/cli)

```shell
brew install tektoncd-cli
```

## Usage

Install [`git-clone`](https://github.com/tektoncd/catalog/tree/main/task/git-clone/0.4) `Task` from the Tekton Catalog

```shell
kubectl apply -f https://raw.githubusercontent.com/tektoncd/catalog/main/task/git-clone/0.4/git-clone.yaml
```

Install [`golang-test`](https://github.com/tektoncd/catalog/tree/main/task/golang-test/0.2) `Task` from the Tekton Catalog

```shell
kubectl apply -f https://raw.githubusercontent.com/tektoncd/catalog/main/task/golang-test/0.2/golang-test.yaml
```

Install [`golang-fuzz`](https://github.com/tektoncd/catalog/tree/main/task/golang-fuzz/0.1) `Task` from the Tekton Catalog

```shell
kubectl apply -f https://raw.githubusercontent.com/tektoncd/catalog/main/task/golang-fuzz/0.1/golang-fuzz.yaml
```

Run the `Pipeline` and show the logs

```shell
tkn pipeline start --showlog \
-f tekton/pipeline.yaml \
-p package=github.com/jerop/fuzzton \
-w name=workarea,volumeClaimTemplateFile=tekton/workspace.yaml
```
