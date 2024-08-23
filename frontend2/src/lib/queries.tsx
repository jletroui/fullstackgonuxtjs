import { API_BASE_URL } from "./config"
import { Signal, signal } from '@preact/signals'
import { useEffect } from 'preact/hooks';

// This module implements a poor man's SWR
// If we need more flexibility, especially more granularity within stored data for reactivity, this is probably a good alternative:
// https://github.com/nanostores/query + https://github.com/nanostores/preact

export interface QueryResponse<T> {
    // Signals are maybe a bit overkill, since it is safe and performant to use the query responses in multiple component. It is unlikely
    // we really need to pass one of these down a component tree. But, it might be convenient for large components.
    data: Signal<T | undefined>
    error: Signal<any> // Todo: standardized app error
    isLoading: Signal<boolean>
    fetchCount: Signal<number>
    refresh: () => void,
}

enum QueryStatus {
    idle, loading, done
}

class Query<Resp, Val> implements QueryResponse<Val> {
    readonly data = signal<Val | undefined>(undefined)
    readonly error = signal<any>(null)
    readonly isLoading = signal(false)
    readonly fetchCount = signal(1)
    private status = QueryStatus.idle
    private readonly url: string
    private readonly transform: (resp: Resp) => Val

    constructor(url: string, transform: (resp: Resp) => Val) {
        this.url = url
        this.transform = transform
    }

    readonly refresh = () => {
        this.status = QueryStatus.idle
        this.fetchCount.value = this.fetchCount.value + 1; // Signals to useEffect() this is dirty and needs to be refetched next render.
    }

    fetch() {
        // We need to guard with this, since multiple components might still hit the same query
        // Indeed: each component using the same query will have its useEffect() fired (line 89)
        if (this.status === QueryStatus.idle) {
            this.status = QueryStatus.loading
            this.error.value = null
            this.isLoading.value = true
            fetch(`${API_BASE_URL}${this.url}`).then(
                async resp => {
                    if (resp.status >= 200 && resp.status < 400) {
                        const jsonResp = await resp.json() as Resp;
                        this.data.value = this.transform(jsonResp)
                        this.isLoading.value = false
                    } else {
                        this.data.value = undefined
                        this.isLoading.value = false
                        this.error.value = this.isJson(resp) ? await resp.json() : await resp.text()
                    }
                    this.isLoading.value = false
                    this.status = QueryStatus.done
                },
                errResp => {
                    this.data.value = undefined
                    this.error.value = errResp
                    this.isLoading.value = false
                    this.status = QueryStatus.done
                }
            )
        }
    }

    private isJson(resp: Response) {
        return resp.headers.get('content-type')?.includes('application/json') ?? false
    }
}

const cache = new Map<string, Query<any, any>>()

// Useful between tests, where we want to test how the query is handling multiple type of responses 
export const clearQueryCache = () => {
    cache.clear()
}

const getQuery = <Resp, Val>(url: string, transform: (resp: Resp) => Val) => {
    if (!cache.has(url)) {
        cache.set(url, new Query(url, transform))
    }
    return cache.get(url) as Query<Resp, Val>
}

export const  useQuery = <Resp, Val>(url: string, transform: (resp: Resp) => Val): QueryResponse<Val> => {
    const query = getQuery(url, transform)
    useEffect(() => query.fetch(), [url, query.fetchCount.value])
    return query;
}

interface TaskCount {
    count: number
}
export const useTaskCount = () => useQuery('/tasks/count', (resp: TaskCount) => resp.count)
