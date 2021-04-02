{{/*
Expand the name of the chart.
*/}}
{{- define "kegger.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "kegger.fullname" -}}
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
Frontend & Backend customizations
*/}}

{{- define "frontend.fullname" -}}
{{- printf "%s-frontend" .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "backend.fullname" -}}
{{- printf "%s-api" .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "frontend-deployment.fullname" -}}
{{- printf "%s-frontend" .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "backend-deployment.fullname" -}}
{{- printf "%s-api" .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "backend-service.fullname" -}}
{{- printf "%s-backend-svc" .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "frontend-service.fullname" -}}
{{- printf "%s-frontend-svc" .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "backend-service.url" -}}
{{- printf "http://%s-backend-svc:4040" .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}


{{/*
Return postgresql service fullname
*/}}
{{- define "kegger.postgresql.fullname" -}}
{{- printf "%s-%s" .Release.Name "postgresql" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Return postgresql host
*/}}
{{- define "kegger.databaseHost" -}}
{{- if .Values.postgresql.enabled }}
    {{- printf "%s" (include "kegger.postgresql.fullname" .) -}}
{{- end -}}
{{- end }}


{{/*
Return postgresql username
*/}}
{{- define "kegger.databaseUser" -}}
{{- if .Values.postgresql.enabled }}
    {{- printf "%s" .Values.postgresql.postgresqlUsername -}}
{{- end -}}
{{- end }}

{{/*
Return postgresql database name
*/}}
{{- define "kegger.databaseName" -}}
{{- if .Values.postgresql.enabled }}
    {{- printf "%s" .Values.postgresql.postgresqlDatabase -}}
{{- end -}}
{{- end }}


{{/*
Return postgresql port
*/}}
{{- define "kegger.databasePort" -}}
{{- if .Values.postgresql.enabled }}
    {{- printf "5432" -}}
{{- end -}}
{{- end }}


{{/*
Return the postgres Secret Name
*/}}
{{- define "kegger.databaseSecretName" -}}
{{- if .Values.postgresql.enabled }}
    {{- printf "%s" (include "kegger.postgresql.fullname" .) -}}
{{- end -}}
{{- end -}}



{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "kegger.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "kegger.labels" -}}
helm.sh/chart: {{ include "kegger.chart" . }}
{{ include "kegger.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "kegger.selectorLabels" -}}
app.kubernetes.io/name: {{ include "kegger.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "kegger.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "kegger.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
