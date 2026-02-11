
// This interface exists to allow us to drop in different
// API implementations. Especially while nox undergoes 
// heavy development. I want this to be a very easy way to
// mock incoming data until I can actually deploy a webserver
// environment.
//
// And this could serve as a good expandable way later on to
// allow for mods or custom backends. Although Im not confident
// that we will ever do such a thing

export interface IBackendApi {
    route: string,
    jwt: string,
    loggedIn: boolean, //temporary to mimick auth

    login(username: string, password: string): Promise<void>,
    refreshTok(): Promise<void>,

    getUser(username: string): Promise<any>,
    updateUser(fn: (u: any) => any): Promise<void>, //runs the callback to update user, then updates in backend eventually
    removeUser(id: string): Promise<void>,
    createUser(user: any): Promise<any>,
}

class MockApi implements IBackendApi {
    route: string = "http://mock.com/";
    jwt: string = "somejwt";
    loggedIn: boolean = true;
    fakeDb: Map<string, any> = new Map<string, any>;

    async login(username: string, password: string): Promise<void> {

    }
    async refreshTok(): Promise<void> {

    }

    async getUser(username: string): Promise<any> {

    }
    async updateUser(fn: (u: any) => any): Promise<void> {

    }
    async removeUser(id: string): Promise<void> {

    }
    async createUser(user: any): Promise<any> {

    }
}
