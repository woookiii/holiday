import { QueryClient } from "@tanstack/query-core";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: false,
      staleTime: 20 * 1000, //ms, same request will be get from cache
    },
    mutations: {
      retry: false,
    },
  },
});

export default queryClient;
