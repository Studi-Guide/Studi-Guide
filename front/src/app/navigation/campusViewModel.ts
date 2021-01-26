import {ICampus} from '../building-objects-if';
import {LatLngLiteral} from 'leaflet';
export class CampusViewModel {

    private readonly campus: ICampus;
    private readonly addressString: string;
    constructor(campus: ICampus) {
        this.campus = campus;
        this.addressString = this.getAddressString(this.campus);
    }

    public get Address() {
        return this.addressString;
    }

    public get Name() {
        return this.campus.Name;
    }

    public get ShortName() {
        return this.campus.ShortName;
    }

    public get LatLng(): LatLngLiteral {
        return {lat: this.campus.Latitude, lng: this.campus.Longitude};
    }

    private getAddressString(campus: ICampus) {
        let addressString = '';
        if (campus.edges.Buildings !== null) {

            const addresses = campus.edges.Buildings.map(x => x.edges.Address);
            const addressGroup = groupBy(addresses, i => i.Street);
            for (const street of Object.keys(addressGroup)) {
                const group = addressGroup[street];
                if (addressString !== '') {
                    addressString += '; ';
                }

                let addressNumber = group[0].Number;
                if (group.length > 1 && group[group.length - 1].Number !== addressNumber) {
                    addressNumber += '-' + group[group.length - 1].Number;
                }

                addressString += street + ' ' + addressNumber + ', ';
                addressString += addressGroup[street][0].PLZ + ' ' + addressGroup[street][0].City;
            }
        }

        return addressString;
    }
}

const groupBy = <T, K extends keyof any>(list: T[], getKey: (item: T) => K) =>
    list.reduce((previous, currentItem) => {
        const group = getKey(currentItem);
        if (!previous[group]) { previous[group] = []; }
        previous[group].push(currentItem);
        return previous;
    }, {} as Record<K, T[]>);

