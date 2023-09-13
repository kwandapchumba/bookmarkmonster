import { apiHost, session } from '../stores/stores';
import type { Session } from '../types/session';
import type { Folder } from '../types/folder';

let apiEndpoint: string;
let userSession: Partial<Session>;

function getApiHost() {
	const unsub = apiHost.subscribe((value) => {
		apiEndpoint = value;
	});

	unsub();
}

function getUserSession() {
	const unsub = session.subscribe((value) => {
		userSession = value;
	});

	unsub();
}
export async function getFolderPath(folderID: string, fetch: any): Promise<Folder[]> {
	getUserSession();

	getApiHost();

	const response = await fetch(
		`${apiEndpoint}/authenticated/collections/getFolderPath/${folderID}`,
		{
			method: 'GET',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				authorization: `Bearer${userSession.AccessToken}`
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer'
		}
	);

	const result = await response.json();

	return result[0];
}
