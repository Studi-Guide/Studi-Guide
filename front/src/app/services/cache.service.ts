import {Injectable} from '@angular/core';
import {EMPTY, Observable} from 'rxjs';
import {catchError, shareReplay} from 'rxjs/operators';
import {HttpClient} from '@angular/common/http';

@Injectable({
    providedIn: 'root'
})
export class CacheService {

    cache = {};

    Get<T>(client: HttpClient,  request: string): Observable<T> {
        if (this.cache[request]) {
            return this.cache[request] as Observable<T>;
        }

        this.cache[request] = client.get<T>(request).pipe(
            shareReplay(1),
            catchError(err => {
                delete this.cache[request];
                return EMPTY;
            }));

        return this.cache[request];
    }
}
