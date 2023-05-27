// Query provider for context using @tanstackquery
"use client";

import {
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";

const queryClient = new QueryClient();

const ReactQueryrovider = ( { children } : { children: React.ReactNode}) => {
    
    return (
        <QueryClientProvider client={ queryClient }>
            { children }
        </QueryClientProvider>
    );
};

export default ReactQueryrovider;
