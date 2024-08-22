import { API_BASE_URL } from "./config"
import { Signal, useSignal } from '@preact/signals'
import { useEffect, useMemo, useState } from 'preact/hooks';

// Poor man's SWR
export const  useQuery = <Resp, Val>(url: string, transform: (resp: Resp) => Val): QueryResponse<Val> => {
    return useMemo(() => {
        const data = useSignal<Val | undefined>()
        const error = useSignal<any>(null)
        const isLoading = useSignal(true)
        const [ fetchCount, setFetchCount ] = useState(1)

        useEffect(() => {
            error.value = null
            isLoading.value = true
            fetch(`${API_BASE_URL}${url}`)
                .then(
                    async resp => {
                        if (resp.status >= 200 && resp.status < 400) {
                            const jsonResp = await resp.json() as Resp;
                            data.value = transform(jsonResp)
                            isLoading.value = false
                        } else {
                            data.value = undefined
                            isLoading.value = false
                            error.value = resp.headers.get('content-type')?.includes('application/json') ? await resp.json() : await resp.text()
                        }
                    },
                    errResp => {
                        data.value = undefined
                        isLoading.value = false
                        error.value = errResp
                    }
                )
        }, [url, fetchCount])

        return {
            data,
            error,
            isLoading,
            refresh: () => setFetchCount(cnt => cnt + 1)
        }
    }, [url])
}

export interface QueryResponse<T> {
    data: Signal<T | undefined>
    error: Signal<any>
    isLoading: Signal<boolean>
    refresh: () => void,
}

interface TaskCount {
    count: number
}

export const useTaskCount = () => useQuery('/tasks/count', (resp: TaskCount) => resp.count)
