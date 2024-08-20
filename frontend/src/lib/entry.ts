// file: src/lib/entry.ts

export type Entry = {
	id: number;
	notebook_id: number;
	author_id: number;
	title: string;
	content: string;
	role: string;
	parent_entry_id?: number | null;
	timestamp: string;
	has_photo: boolean;
	tag_id: number | null;
	children: Entry[];
};

export type LocalEntryType = {
	id: number;
	role: string;
	content: string;
	timestamp?: string;
	parent_entry_id?: number | null;
	tag_id: number | null;
};

type FetchedEntry = {
	id: number;
	notebook_id: number;
	author_id: number;
	parent_entry_id: number | null;
	title: string;
	content: string;
	role: string;
	timestamp: string;
	has_photo: boolean;
	tag_id: number | null;
};

export type NewEntry = {
	title: string;
	content?: string;
	role?: string;
	tag_id?: number | null;
	parent_entry_id?: number | null;
};

export type NewChildEntry = {
	title: string;
	content?: string;
	role?: string;
	tag_id?: number | null;
	parent_entry_id?: number | null;
};

import api from './api';

export function evalTimePassed(timestamp: string) {
	const date = new Date(timestamp);
	const today = new Date();

	const timePassed = Math.floor((today.getTime() - date.getTime()) / 1000);

	if (timePassed < 60) {
		return `${timePassed} second${timePassed === 1 ? '' : 's'} ago`;
	} else if (timePassed < 3600) {
		return `${Math.floor(timePassed / 60)} minute${Math.floor(timePassed / 60) === 1 ? '' : 's'} ago`;
	} else if (timePassed < 86400) {
		return `${Math.floor(timePassed / 3600)} hour${Math.floor(timePassed / 3600) === 1 ? '' : 's'} ago`;
	} else if (timePassed < 604800) {
		return `${Math.floor(timePassed / 86400)} day${Math.floor(timePassed / 86400) === 1 ? '' : 's'} ago`;
	} else {
		return `${Math.floor(timePassed / 604800)} week${Math.floor(timePassed / 604800) === 1 ? '' : 's'} ago`;
	}
}

export async function fetchEntries(params: { sessionId: string; slug: string }): Promise<Entry[]> {
	const res = await api.get(`/api/notebooks/${params.slug}/entries`, {
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${params.sessionId}`
		}
	});

	if (res.status === 200) {
		const fetchedEntries = res.data as FetchedEntry[];
		const entriesMap = new Map<number, Entry>();

		// First pass: Create Entry objects and store them in the map

		if (fetchedEntries && fetchedEntries.length > 0) {
			fetchedEntries?.forEach((fe) => {
				entriesMap.set(fe.id, {
					...fe,
					children: []
				} as Entry);
			});

			// Second pass: Organize entries into a tree structure
			const rootEntries: Entry[] = [];
			fetchedEntries?.forEach((fe) => {
				const entry = entriesMap.get(fe.id)!;
				if (fe.parent_entry_id === null) {
					rootEntries.push(entry);
				} else {
					const parentEntry = entriesMap.get(fe.parent_entry_id);
					if (parentEntry) {
						parentEntry.children.push(entry);
					}
				}
			});

			rootEntries?.forEach((e) => {
				e.children.sort((a, b) => a.timestamp.localeCompare(b.timestamp));
			});

			return rootEntries;
		} else {
			return [];
		}
	} else {
		return [];
	}
}

export async function addEntry(
	sessionId: string,
	slug: string,
	entry: NewEntry
): Promise<Entry | null> {
	const res = await api.post(`/api/notebooks/${slug}/entries`, entry, {
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${sessionId}`
		}
	});

	if (res.status === 200) {
		return res.data;
	} else {
		return null;
	}
}

export async function addChildEntry(
	sessionId: string,
	slug: string,
	entryId: number,
	childEntry: NewChildEntry
): Promise<Entry | null> {
	console.log(entryId, childEntry);
	try {
		const res = await api.post(`/api/notebooks/${slug}/entries/${entryId}/children`, childEntry, {
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${sessionId}`
			}
		});

		if (res.status === 200) {
			return res.data;
		} else {
			console.log(res.data);
			return null;
		}
	} catch (error) {
		// @ts-expect-error asd
		console.log(error.response.data);
		return null;
	}
}

export async function fetchChildEntries(
	sessionId: string,
	slug: string,
	entryId: number
): Promise<Entry[] | null> {
	const res = await api.get(`/api/notebooks/${slug}/entries/${entryId}/children`, {
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${sessionId}`
		}
	});

	if (res.status === 200) {
		return res.data;
	} else {
		return null;
	}
}

export async function updateChildEntry(
	sessionId: string,
	slug: string,
	entryId: number,
	childEntryId: number,
	childEntry: NewChildEntry
): Promise<Entry | null> {
	console.log(childEntryId, childEntry);
	try {
		const res = await api.patch(
			`/api/notebooks/${slug}/entries/${entryId}/children/${childEntryId}`,
			childEntry,
			{
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${sessionId}`
				}
			}
		);

		if (res.status === 200) {
			return res.data;
		} else {
			console.log(res.data);
			return null;
		}
	} catch (error) {
		// @ts-expect-error asd
		console.log(error.response.data);
		return null;
	}
}

export async function updateEntry(
	sessionId: string,
	slug: string,
	entryId: number,
	entry: NewEntry
): Promise<Entry | null> {
	console.log(entryId, entry);
	try {
		const res = await api.patch(`/api/notebooks/${slug}/entries/${entryId}`, entry, {
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${sessionId}`
			}
		});

		if (res.status === 200) {
			return res.data;
		} else {
			console.log(res.data);
			return null;
		}
	} catch (error) {
		// @ts-expect-error asd
		console.log(error.response.data);
		return null;
	}
}

export async function deleteEntry(
	sessionId: string,
	slug: string,
	entryId: number
): Promise<{ id: number } | null> {
	try {
		const res = await api.delete(`/api/notebooks/${slug}/entries/${entryId}`, {
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${sessionId}`
			}
		});

		if (res.status === 200) {
			return res.data;
		} else {
			console.log(res.data);
			return null;
		}
	} catch (error) {
		// @ts-expect-error asd
		console.log(error.response.data);
		return null;
	}
}

export function convertEntriesToLocal(entry: Entry): LocalEntryType[] {
	const result: LocalEntryType[] = [];

	function processEntry(entry: Entry) {
		result.push({
			id: entry.id,
			role: entry.role,
			content: entry.title,
			timestamp: entry.timestamp,
			parent_entry_id: entry.parent_entry_id,
			tag_id: entry.tag_id
		});

		// Process children recursively
		if (entry.children && entry.children.length > 0) {
			entry.children.forEach(processEntry);
		}
	}

	processEntry(entry);

	return result;
}
