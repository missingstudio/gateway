"use client";

import { ColumnDef } from "@tanstack/react-table";

import { Badge } from "@missingstudio/ui/badge";
import { Checkbox } from "@missingstudio/ui/checkbox";

import { DataTableColumnHeader } from "~/app/(llm)/playground/components/datatable/columnheader";
import { DataTableRowActions } from "~/app/(llm)/playground/components/datatable/rowactions";

export const columns: ColumnDef<any>[] = [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && "indeterminate")
        }
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
        className="translate-y-[2px]"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
        className="translate-y-[2px]"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },

  {
    accessorKey: "model",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Model" />
    ),
    cell: ({ row }) => {
      return (
        <div className="flex space-x-2">
          <Badge variant="outline">{row.original.provider}</Badge>
          <span className="max-w-[150px] truncate font-medium">
            {row.original.model}
          </span>
        </div>
      );
    },
    enableSorting: false,
  },

  {
    accessorKey: "prompt_tokens",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Prompt Tokens" />
    ),
    cell: ({ row }) => {
      return (
        <div className="flex space-x-2">
          <span className="max-w-[50px] truncate font-medium">
            {row.original.prompt_tokens}
          </span>
        </div>
      );
    },
  },
  {
    accessorKey: "completion_tokens",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Completion Tokens" />
    ),
    cell: ({ row }) => {
      return (
        <div className="flex space-x-2">
          <span className="max-w-[50px] truncate font-medium">
            {row.original.completion_tokens}
          </span>
        </div>
      );
    },
  },
  {
    accessorKey: "total_tokens",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Total Tokens" />
    ),
    cell: ({ row }) => {
      return (
        <div className="flex space-x-2">
          <span className="max-w-[50px] truncate font-medium">
            {row.original.total_tokens}
          </span>
        </div>
      );
    },
  },
  {
    accessorKey: "latency",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Latency" />
    ),
    cell: ({ row }) => {
      return (
        <div className="flex space-x-2">
          <span className="max-w-[100px] truncate font-medium">
            {row.original.latency}
          </span>
        </div>
      );
    },
  },

  {
    id: "actions",
    cell: ({ row }) => <DataTableRowActions row={row} />,
  },
];
