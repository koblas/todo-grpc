interface HasId {
  id: string;
}

interface Event<T extends HasId> {
  object_id: string;
  action: string;
  body: T | null;
}

export function sharedEventCrud<Item extends HasId>(old: Item[] | undefined, event: Event<Item>): Item[] {
  const items: Item[] = old ?? [];

  if (event.action === "delete") {
    return items.filter(({ id }) => id !== event.object_id);
  }

  if (!event.body) {
    throw new Error(`Missing event body: ${event.action}`);
  }
  // Moving to local variable to make TS happy with typing
  const item = event.body;

  if (event.action === "create") {
    // If this already is present, then ignore
    if (items.some(({ id }) => id === item.id)) {
      return items;
    }
    return items.concat([item]);
  }

  if (event.action === "update") {
    return items.map((value) => {
      if (value.id === item.id) {
        return item;
      }

      return value;
    });
  }

  throw new Error(`Unknown event received: ${event.action}`);
}
