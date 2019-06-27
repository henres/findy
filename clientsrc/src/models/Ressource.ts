import { Tag } from '@/models/Tag';

export class Ressource {
    public id: number
    public kind: string
    public description: string
    public location: string
    public tags: Tag[]

    constructor(
        id: number,
        kind: string,
        description: string,
        location: string,
        tags: Tag[],
    ) {
        this.id = id
        this.kind = kind
        this.description = description
        this.location = location
        this.tags = tags
    }
}
