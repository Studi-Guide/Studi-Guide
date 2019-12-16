# backend ----------------------
FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build-env
WORKDIR /src/main/back/StudiGuideAppService
#COPY . .
#RUN dotnet restore

#COPY . ./
#RUN dotnet publish StudiGuideAppService/StudiGuideAppService.csproj -c Release -o out

# Build runtime image
#FROM mcr.microsoft.com/dotnet/core/aspnet:3.1
#WORKDIR /src/main/back
#COPY --from=build-env src/main/back/out .
#ENTRYPOINT ["dotnet", "StudiGuideAppService.dll"]

# ---------------------------------

# Copy csproj and restore as distinct layers
COPY *.csproj ./
RUN dotnet restore

# Copy everything else and build
COPY . ./
RUN dotnet publish -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/core/aspnet:3.1
WORKDIR /src/main/back/StudiGuideAppService
COPY --from=build-env /app/out .
ENTRYPOINT ["dotnet", "StudiGuideAppService.dll"]