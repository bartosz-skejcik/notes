// notes=# select * from sticky_notes;
//  id | notebook_id | category_id | author_id | title | content | created_at
// ----+-------------+-------------+-----------+-------+---------+------------
// (0 rows)

import api from '$lib/api';

export interface StickyNote {
  id: number;
  notebook_id: number;
  category_id: number;
  author_id: number;
  title: string;
  content: string;
  created_at: string;
}

export async function getStickyNotes(notebookId: string, sessionId: string): Promise<StickyNote[]> {
  const res = await api.get(`/api/notebooks/${notebookId}/categories/sticky-notes`, {
    headers: {
      Authorization: `Bearer ${sessionId}`
    }
  });

  if (res.status === 200) {
    return res.data;
  } else {
    console.error(res.data);
  }

  return [];
}

export async function getStickyNotesForUserCategory(
  notebookId: string,
  categoryId: string,
  sessionId: string
): Promise<StickyNote[]> {
  const res = await api.get(`/api/notebooks/${notebookId}/categories/${categoryId}/sticky-notes`, {
    headers: {
      Authorization: `Bearer ${sessionId}`
    }
  });

  if (res.status === 200) {
    return res.data;
  } else {
    console.error(res.data);
  }

  return [];
}

export async function createStickyNote(
  notebookId: string,
  categoryId: string,
  sessionId: string,
  title: string,
  content: string
): Promise<{ id: number } | null> {
  const res = await api.post(
    `/api/notebooks/${notebookId}/categories/${categoryId}/sticky-notes`,
    {
      title,
      content
    },
    {
      headers: {
        Authorization: `Bearer ${sessionId}`,
        'Content-Type': 'application/json'
      }
    }
  );

  if (res.status === 200) {
    return res.data;
  } else {
    console.error(res.data);
  }

  return null;
}

export async function updateStickyNote(
  notebookId: string,
  categoryId: string,
  stickyNoteId: number,
  sessionId: string,
  title: string,
  content: string
): Promise<{ id: number } | null> {
  const res = await api.patch(
    `/api/notebooks/${notebookId}/categories/${categoryId}/sticky-notes/${stickyNoteId}`,
    {
      title,
      content
    },
    {
      headers: {
        Authorization: `Bearer ${sessionId}`,
        'Content-Type': 'application/json'
      }
    }
  );

  if (res.status === 200) {
    return res.data;
  } else {
    console.error(res.data);
  }

  return null;
}
