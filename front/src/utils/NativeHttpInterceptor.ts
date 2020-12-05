import {Injectable, Injector} from '@angular/core';
import {HttpInterceptor, HttpRequest, HttpHandler, HttpEvent, HttpResponse, HttpHeaders} from '@angular/common/http';
import { Observable, from } from 'rxjs';
import { Platform } from '@ionic/angular';
import {HTTP, HTTPResponse} from '@ionic-native/http/ngx';

type HttpMethod = 'get' | 'post' | 'put' | 'patch' | 'head' | 'delete' | 'upload' | 'download';

@Injectable()
export class NativeHttpInterceptor implements HttpInterceptor {
    private nativeHttp: HTTP;

    constructor( private platform: Platform) {
    }

    public intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        if (this.platform.is('hybrid') && !this.nativeHttp) {
            this.nativeHttp = new HTTP();
        }

        if (!this.platform.is('hybrid')) { return next.handle(request); }
        return from(this.handleNativeRequest(request));
    }

    private async handleNativeRequest(request: HttpRequest<any>): Promise<HttpResponse<any>> {
        const headerKeys = request.headers.keys();
        const header = {};

        headerKeys.forEach((key) => {
            header[key] = request.headers.get(key);
        });

        try {
            await this.platform.ready();
            const reqMethod = request.method.toLowerCase() as HttpMethod;

            const nativeHttpResponse = await this.nativeHttp.sendRequest(request.url, {
                method: reqMethod,
                data: request.body,
                headers: header,
                serializer: 'json',
            });

            let reqbody;

            try {
                if (!request.responseType || request.responseType === 'json'){
                    reqbody = JSON.parse(nativeHttpResponse.data);
                } else {
                   if (request.responseType === 'text'){
                       reqbody = nativeHttpResponse.data;
                   }
                }

            } catch (error) {
                console.log(error);
                reqbody = nativeHttpResponse.data;
            }

            const reqHeader = new HttpHeaders(nativeHttpResponse.headers);
            const response = new HttpResponse({
                body: reqbody,
                status: nativeHttpResponse.status,
                headers: reqHeader,
                url: nativeHttpResponse.url,
            });

            return Promise.resolve(response);
        } catch (error) {
            if (!error.status) { return Promise.reject(error); }

            const response = new HttpResponse({
                body: JSON.parse(error.error),
                status: error.status,
                headers: error.headers,
                url: error.url,
            });

            return Promise.reject(response);
        }
    }
}