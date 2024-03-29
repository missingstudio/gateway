"use client";

import { PopoverProps } from "@radix-ui/react-popover";
import { Settings } from "lucide-react";
import React, { useEffect } from "react";

import { Button } from "@missingstudio/ui/button";
import { Input } from "@missingstudio/ui/input";
import { Label } from "@missingstudio/ui/label";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@missingstudio/ui/popover";
import { Slider } from "@missingstudio/ui/slider";
import {
  ParameterType,
  useParameterFetch,
} from "~/app/(llm)/playground/hooks/useParameterFetch";
import { useStore } from "~/app/(llm)/playground/store";

interface ModelSelectorProps extends PopoverProps {}

export default function ModelSelector(props: ModelSelectorProps) {
  const [open, setOpen] = React.useState(false);
  const { parameters } = useParameterFetch();
  const { setParameter, getParameter } = useStore();

  useEffect(() => {
    for (const param of parameters) setParameter(param.id, param.default);
  }, [parameters, setParameter]);

  return (
    <div className="flex items-center gap-2">
      <Popover>
        <PopoverTrigger asChild>
          <Button variant="outline">
            <Settings className="-m-1 h-4 w-4" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="">
          <div className="flex flex-col gap-4">
            <div className="space-y-2">
              <h4 className="font-medium leading-none">Parameters</h4>
            </div>

            {parameters.map((param: ParameterType) => {
              if (param.type === "float" || param.type === "int")
                return (
                  <div key={param.id} className="flex flex-col gap-4">
                    <div className="flex items-center justify-between gap-4">
                      <Label htmlFor="width">{param.name}</Label>
                      <Input
                        id={param.id}
                        type="number"
                        value={getParameter(param.id)}
                        className="h-8 w-1/3"
                        onChange={(e) => {
                          setParameter(param.id, parseFloat(e.target.value));
                        }}
                      />
                    </div>
                    <Slider
                      id={param.id}
                      value={[getParameter(param.id) as number]}
                      min={param.min}
                      max={param.max}
                      step={param.step}
                      onValueChange={(value) =>
                        setParameter(param.id, value[0])
                      }
                    />
                  </div>
                );
            })}
          </div>
        </PopoverContent>
      </Popover>
    </div>
  );
}
