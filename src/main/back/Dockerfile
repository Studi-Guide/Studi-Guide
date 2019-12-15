FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build-env

COPY . /app
RUN dotnet publish StudiGuideAppService/StudiGuideAppService.csproj -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/core/aspnet:3.1
COPY --from=build-env /app/out .
ENTRYPOINT ["dotnet", "app/StudiGuideAppService.dll"]
