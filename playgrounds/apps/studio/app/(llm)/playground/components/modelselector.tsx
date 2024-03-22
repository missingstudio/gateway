"use client";
import { CaretSortIcon } from "@radix-ui/react-icons";
import { PopoverProps } from "@radix-ui/react-popover";
import React from "react";

import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@missingstudio/ui/popover";

import { Button } from "@missingstudio/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandList,
} from "@missingstudio/ui/command";
import { Label } from "@missingstudio/ui/label";
import { ModelItem } from "~/app/(llm)/playground/components/modelitem";
import { useModelFetch } from "~/app/(llm)/playground/hooks/useModelFetch";
import { useStore } from "~/app/(llm)/playground/store";

interface ModelSelectorProps extends PopoverProps {}

export default function ModelSelector(props: ModelSelectorProps) {
  const [open, setOpen] = React.useState(false);
  const [isFineTuning, setIsFineTuning] = React.useState(false);
  const { providers } = useModelFetch(isFineTuning);
  const { model, setModel, setProvider } = useStore();

  const toggleFineTuning = () => setIsFineTuning(!isFineTuning);

  return (
    <div className="flex items-center gap-2">
      <Label htmlFor="model">Model: </Label>
      <Button variant="outline" onClick={toggleFineTuning}>
        {isFineTuning ? 'Select for Fine-Tuning' : 'Select Model'}
      </Button>
      <Popover open={open} onOpenChange={setOpen} {...props}>
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            aria-label="Select a model"
            className="w-full justify-between"
          >
            {model ? model : "Select a model..."}
            <CaretSortIcon className="ml-2 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent align="end" className="w-[250px] p-0">
          <Command loop>
            <CommandList className="h-[var(--cmdk-list-height)] max-h-[400px]">
              <CommandInput placeholder="Search Models..." />
              <CommandEmpty>No Models found.</CommandEmpty>
              {providers.map((provider) => (
                <CommandGroup key={provider.name} heading={provider.name}>
                  {provider.models.map((singleModel, index) => (
                    <ModelItem
                      key={`${index}_${provider.name}_${singleModel.value}`}
                      id={`${singleModel.value}`}
                      isSelected={model === singleModel.value}
                      onSelect={() => {
                        setProvider(provider.name);
                        setModel(singleModel.value);
                        setOpen(false);
                      }}
                    />
                  ))}
                </CommandGroup>
              ))}
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>
    </div>
  );
}
