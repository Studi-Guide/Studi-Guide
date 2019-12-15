FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build-env

RUN dotnet publish StudiGuideAppService/StudiGuideAppService.csproj -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/core/aspnet:3.1
COPY --from=build-env /out .
ENTRYPOINT ["dotnet", "StudiGuideAppService.dll"]
