import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Env} from '../../environments/environment';
import {forEach} from '@angular-devkit/schematics';
import {CalenderRequestData, MoodleToken} from '../moodle-objects-if';

@Injectable({
    providedIn: 'root'
})
export class MoodleService {
    moodleUrl:string;// = SERVER_URL // "https://studi-guide.azurewebsites.net"; // for development: http://localhost:8090
    constructor(private httpClient : HttpClient, private env : Env) {
        this.moodleUrl = 'https://moodle3.de';
    }

    getLoginToken(user: string, pw: string) {
        const url =  this.moodleUrl + '/login/token.php' + this.generateUrlQuery({
            username : user,
            password : pw,
            service : 'moodle_mobile_app'
        });
        return this.httpClient.get<MoodleToken>(url);
    }

    getCalenderEventsMonth(token: MoodleToken) {
        return this.moodleRequest<CalenderRequestData>(token, 'core_calendar_get_calendar_monthly_view', null);
    }

    getCalenderEventsWeek(token: MoodleToken) {
        return this.moodleRequest<CalenderRequestData>(token, 'core_calendar_get_calendar_upcoming_view', null);
    }

    public containsEvents(responseObj: any) : boolean {
        return responseObj.events != null || responseObj.events !== undefined
    }

    public containsToken(responseObj: any) : boolean {
        return responseObj.token != null || responseObj.token !== undefined
    }

    private moodleRequest<T>(token: MoodleToken, restFunction: string, parameters: ParameterMap) {
        const url = this.moodleUrl + '/webservice/rest/server.php' +  this.generateUrlQuery({
            wstoken: token.token,
            wsfunction: restFunction,
            moodlewsrestformat: 'json'
        });

        return this.httpClient.get<T>(url);
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
