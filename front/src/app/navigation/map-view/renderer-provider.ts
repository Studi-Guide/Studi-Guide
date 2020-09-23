import {ILocation, IMapItem} from '../../building-objects-if';
import {MapItemRendererCanvas} from './map-item-renderer.canvas';
import {LocationRendererCanvas} from './location-renderer.canvas';
import {IReceivedRoute} from '../../route-objects-if';
import {RouteRendererCanvas} from './route-renderer.canvas';


export class RendererProvider {
    public static GetMapItemRendererCanvas(...mapItem:IMapItem[]) : MapItemRendererCanvas[] {
        const renderer:MapItemRendererCanvas[] = [];
        for (const m of mapItem) {
            renderer.push(new MapItemRendererCanvas(m));
        }
        return renderer;
    }

    public static GetLocationRendererCanvas(...location:ILocation[]) : LocationRendererCanvas[] {
        const renderer:LocationRendererCanvas[] = [];
        for (const l of location) {
            renderer.push(new LocationRendererCanvas(l));
        }
        return renderer;
    }

    public static GetRouteRendererCanvas(...route:IReceivedRoute[]) : RouteRendererCanvas[] {
        const renderer:RouteRendererCanvas[] = [];
        for (const r of route) {
            renderer.push(new RouteRendererCanvas(r));
        }
        return renderer;
    }
}