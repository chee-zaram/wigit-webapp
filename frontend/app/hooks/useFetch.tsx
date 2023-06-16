// data fetching hook
"use client";

import { useState, useEffect } from 'react';

//put optional params
export default function useFetch(url: string) {
    const [data, setData] = useState<any>(null);
    useEffect(() => {
      fetch(url)
        .then(res => res.json())
        .then((dta) => {
          setData(dta)
        })
    }, [])
    return data
  }


    