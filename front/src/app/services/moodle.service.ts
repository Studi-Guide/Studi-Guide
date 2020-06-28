import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Env} from '../../environments/environment';
import {forEach} from '@angular-devkit/schematics';
import {MoodleToken} from '../moodle-objects-if';

@Injectable({
    providedIn: 'root'
})
export class MoodleService {
    baseUrl:string;// = SERVER_URL // "https://studi-guide.azurewebsites.net"; // for development: http://localhost:8090
    constructor(private httpClient : HttpClient, private env : Env) {
        this.baseUrl = 'https://www.moodle3.de';
    }

    getLoginToken(user: string, pw: string) {
        const parameters: { [key: string] : string; } = {
            username : user,
            password : pw,
            service : 'moodle_mobile_app'
        };

        const url =  this.baseUrl + '/login/token.php' + this.generateUrlQuery(parameters);
        return this.httpClient.get<MoodleToken>(url);
    }

    private generateUrlQuery(parameters: ParameterMap) {
        if (parameters == null || Object.keys(parameters).length === 0) {
            return '';
        }
        let urlParams = '?';
        for (const key in parameters) {
            if (parameters.hasOwnProperty(key)) {
                if (urlParams !== '?') {
                    urlParams +='&';
                }
                const value = parameters[key];
                urlParams += key + '=' + value;
            }
        }
        return urlParams;
    }
}

interface ParameterMap{
    [key: string]: string;
}
