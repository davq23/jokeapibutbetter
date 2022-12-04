import { Service } from './service';

export default class ConfigService extends Service {
    public getLanguages() {
        return this.sendRequest('config/languages', 'GET', new Headers(), null);
    }
}
