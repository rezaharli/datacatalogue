export function newTableObject() {
    return {
        isLoading: true,
        source: [],
        display: [],
        pagination: {},
        totalItems: 0,
        colWidth: {},
        filterTypes: {},
        reset: true,
    }
}