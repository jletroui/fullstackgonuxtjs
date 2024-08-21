import { API_BASE_URL } from "./config"
import { Signal, useSignal } from '@preact/signals'
import { useEffect, useState } from 'preact/hooks';

const fetcher = async <T>(path: string): Promise<T> => {
    const resp = await fetch(`${API_BASE_URL}${path}`)
    return await resp.json()
}

// Poor man's SWR
export const  useQuery = <Resp, Val>(url: string, defaultVal: Val, transform: (resp: Resp) => Val): QueryResponse<Val> => {
    const data = useSignal(defaultVal)
    const error = useSignal(null)
    const isLoading = useSignal(true)
    const [ loadingCount, setLoadingCount ] = useState(1)

    useEffect(() => {
        error.value = null
        isLoading.value = true
        fetcher<Resp>(url)
            .then(resp => {
                data.value = transform(resp)
                isLoading.value = false
            })
            .catch(errResp => {
                data.value = defaultVal
                isLoading.value = false
                error.value = errResp
            })
            
    }, [loadingCount])

    return {
        data,
        error,
        isLoading,
        invalidate: () => setLoadingCount(cnt => cnt + 1)
    };
}

export interface QueryResponse<T> {
    data: Signal<T>
    error: Signal<any>
    isLoading: Signal<boolean>
    invalidate: () => void,
}

interface TaskCount {
    count: number
}

export const useTaskCount = () => useQuery<TaskCount, number>('/tasks/count', -1, resp => resp.count)
