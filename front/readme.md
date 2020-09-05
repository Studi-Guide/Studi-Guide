# Ionic Frontend

## Installation
 - `npm install -g @ionic/cli`
 
## Run
 - `ionic serve`
 
## Build
 - `ionic build --prod --engine=browser --localize`
 
## Documentation
 - [Ionic Framework](https://ionicframework.com/docs/installation/cli)


## Internationalization i18n

 - Extract Messages: `ng xi18n`
 - For each language copy `messages.xlf` to `src/locale/messages.<locale identifier>.xlf`
 - You may use a diff viewer to copy previous translations to the copied file
 - Build web app with localization using the command above