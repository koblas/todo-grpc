{{/*
Expand the name of the chart.
*/}}
{{- define "charts.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "charts.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "charts.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "charts.labels" -}}
helm.sh/chart: {{ include "charts.chart" . }}
{{ include "charts.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "charts.selectorLabels" -}}
app.kubernetes.io/name: {{ include "charts.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "charts.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "charts.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Support envFron
*/}}
{{- define "shared.envFrom" -}}
{{- if .Values.envFrom }}
{{- if .Values.envFrom.normal }}
{{- range $key, $value := .Values.envFrom.normal }}
- name: {{ $key }}
  valueFrom:
    configMapKeyRef:
      key: {{ $key }}
      name: {{ $value }}
{{- end }}
{{- end }}
{{- if .Values.envFrom.secret }}
{{- range $key, $value := .Values.envFrom.secret }}
{{- if contains "|" $value }}
- name: {{ $key }}
  valueFrom:
    secretKeyRef:
      key: {{ regexReplaceAll "(.*)\\|(.*)" $value "${2}" }}
      name: {{ regexReplaceAll "(.*)\\|(.*)" $value "${1}" }}
{{- else }}
- name: {{ $key }}
  valueFrom:
    secretKeyRef:
      key: {{ $key }}
      name: {{ $value }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
