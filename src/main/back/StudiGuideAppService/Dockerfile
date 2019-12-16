# backend ----------------------
FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build-env
WORKDIR /src/main/back/StudiGuideAppService
COPY . .
RUN dotnet restore ./StudiGuideAppService.sln

COPY . .

RUN dotnet publish ./StudiGuideAppService.csproj -c Release -o /out

# Build runtime image
FROM mcr.microsoft.com/dotnet/core/aspnet:3.1
WORKDIR /src/main/back/StudiGuideAppService
COPY --from=build-env /out
ENTRYPOINT ["dotnet", "StudiGuideAppService.dll"]

# ---------------------------------
