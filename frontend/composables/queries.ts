interface TaskCountBody {
    count: number
}

export const useTaskCountQuery = async () => {
    const config = useRuntimeConfig();

    const { status, data: cnt } = await useLazyFetch('/tasks/count', {
        baseURL: config.public.apiBaseUrl,
        transform: (body: TaskCountBody) => {
            return body.count
        },
    });

    return { status, cnt };
}
