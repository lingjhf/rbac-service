export interface SplitterOptions {
  height: number
  items?: SplitterItem[]
}

export interface SplitterItem {
  height: number
  minHeight: number
}

class Splitter {
  _items: SplitterItem[]
  _height: number
  constructor({ height, items }: SplitterOptions) {
    this._items = items ?? []
    this._height = height
  }
  addItem(item: SplitterItem): number {
    let currentHeight = item.height
    for (let i = this._items.length - 1; i >= 0; i--) {
      const prevItem = this._items[i]
      const prevItemHeight = prevItem.height - currentHeight
      if (prevItemHeight > prevItem.minHeight) {
        prevItem.height = prevItemHeight
        break
      }
      prevItem.height = prevItem.minHeight
      currentHeight = prevItem.minHeight - prevItemHeight
    }
    return this._items.push(item)
  }
  setItem(index: number, item: Partial<SplitterItem>) {
    const currentItem = this._items[index]
    if (!currentItem) {
      throw new Error('not found index')
    }
    if (item.height === undefined && item.minHeight == undefined) {
      return
    }
    const newItem = {
      height: item.height ?? currentItem.height,
      minHeight: item.minHeight ?? currentItem.minHeight,
    }
    if (newItem.height - currentItem.height > 0) {
    } else {
    }   
  }

}
