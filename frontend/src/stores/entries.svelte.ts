import {
	fetchEntries,
	addEntry,
	addChildEntry,
	updateChildEntry,
	updateEntry,
	deleteEntry,
	type Entry,
	type NewEntry,
	type NewChildEntry
} from '$lib/entry';

let entries: Entry[] = $state([]);

export function useEntriesStore() {
	async function getEntries(sessionId: string, slug: string): Promise<Entry[]> {
		const fetchedEntries = await fetchEntries({ sessionId, slug });
		entries = fetchedEntries;
		return entries;
	}

	async function addNewEntry(
		sessionId: string,
		slug: string,
		entry: NewEntry
	): Promise<Entry | null> {
		const newEntry = await addEntry(sessionId, slug, entry);
		if (newEntry) {
			entries = [newEntry, ...entries];
		}
		return newEntry;
	}

	async function addNewChildEntry(
		sessionId: string,
		slug: string,
		parentId: number,
		childEntry: NewChildEntry
	): Promise<Entry | null> {
		const newChildEntry = await addChildEntry(sessionId, slug, parentId, childEntry);
		if (newChildEntry) {
			updateEntryInStore(parentId, (entry) => {
				entry.children = [...entry.children, newChildEntry];
				return entry;
			});
		}
		return newChildEntry;
	}

	async function updateExistingEntry(
		sessionId: string,
		slug: string,
		entryId: number,
		updatedEntry: NewEntry
	): Promise<Entry | null> {
		const updated = await updateEntry(sessionId, slug, entryId, updatedEntry);
		if (updated) {
			updateEntryInStore(entryId, () => updated);
		}
		return updated;
	}

	async function updateExistingChildEntry(
		sessionId: string,
		slug: string,
		parentId: number,
		childId: number,
		updatedChildEntry: NewChildEntry
	): Promise<Entry | null> {
		const updated = await updateChildEntry(sessionId, slug, parentId, childId, updatedChildEntry);
		if (updated) {
			updateEntryInStore(parentId, (entry) => {
				entry.children = entry.children.map((child) => (child.id === childId ? updated : child));
				return entry;
			});
		}
		return updated;
	}

	async function deleteExistingEntry(
		sessionId: string,
		slug: string,
		entryId: number
	): Promise<boolean> {
		const deleted = await deleteEntry(sessionId, slug, entryId);
		if (deleted) {
			entries = entries.filter((entry) => entry.id !== entryId);
			return true;
		}
		return false;
	}

	async function updateEntryTag(
		sessionId: string,
		slug: string,
		entryId: number,
		tagId: number,
		title: string
	) {
		await updateExistingEntry(sessionId, slug, entryId, { title, tag_id: tagId });
	}

	async function streamAIResponse(sessionId: string, slug: string, entryId: number) {
		// Implement AI response streaming logic here
		// This is a placeholder and should be replaced with actual implementation
		console.log(`Streaming AI response for entry ${entryId}`);
	}

	function updateEntryInStore(entryId: number, updateFn: (entry: Entry) => Entry) {
		entries = entries.map((entry) => {
			if (entry.id === entryId) {
				return updateFn(entry);
			}
			return entry;
		});
	}

	return {
		get entries() {
			return entries;
		},
		getEntries,
		addEntry: addNewEntry,
		addChildEntry: addNewChildEntry,
		updateEntry: updateExistingEntry,
		updateChildEntry: updateExistingChildEntry,
		deleteEntry: deleteExistingEntry,
		updateEntryTag,
		streamAIResponse
	};
}
