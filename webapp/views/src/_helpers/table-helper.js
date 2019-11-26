export function newTableObject() {
    return {
        isLoading: false,
        source: [],
        display: [],
        pagination: {},
        totalItems: 0,
        colWidth: {},
        filterTypes: {},
        reset: true,
    }
}