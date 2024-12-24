{{- define "distributed-crawler.labels" -}}
helm.sh/chart: {{ include "distributed-crawler.chart" . }}
app.kubernetes.io/name: {{ include "distributed-crawler.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "distributed-crawler.chart" -}}
{{ .Chart.Name }}-{{ .Chart.Version }}
{{- end -}}

{{- define "distributed-crawler.name" -}}
{{ .Chart.Name }}
{{- end -}}
