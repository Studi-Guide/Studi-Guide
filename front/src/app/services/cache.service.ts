import {Injectable} from '@angular/core';
import {EMPTY, Observable} from 'rxjs';
import {catchError, shareReplay} from 'rxjs/operators';
import {HttpClient} from '@angular/common/http';

@Injectable({
    providedIn: 'root'
})
export class CacheService {

    cache = {};

    Get<T>(client: HttpClient,  request: string, logOnError: boolean = true): Observable<T> {
        if (this.cache[request]) {
            return this.cache[request] as Observable<T>;
        }

        console.log('CacheService: Request received ' + request);
        this.cache[request] = client.get<T>(request).pipe(
            shareReplay(1),
            catchError(err => {
                if (logOnError) {
                    console.log(err);
                }

                delete this.cache[request];
                return EMPTY;
            }));

        return this.cache[request];
    }
}
