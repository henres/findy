import axios, { AxiosResponse } from 'axios'
import { Ressource } from '@/models/Ressource'

const api = 'http://localhost:3000/api/ressources'

export class RessourceService {
    public static deleteRessource(ressource: Ressource): any {
        return axios.delete(`${api}/${ressource.id}`)
    }

    public static getRessources(): any {
        return axios.get(`${api}`)
    }

    public static addRessource(ressource: Ressource): Promise<AxiosResponse<any>> {
        return axios.post(`${api}`, { ressource })
    }

    public static updateRessource(ressource: Ressource): Promise<AxiosResponse<any>> {
        return axios.put(`${api}/${ressource.id}`, { ressource })
    }
}
