// @generated by protoc-gen-es v1.6.0 with parameter "target=ts,import_extension="
// @generated from file prompt/v1/models.proto (package prompt.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Struct } from "@bufbuild/protobuf";

/**
 * @generated from message prompt.v1.Prompt
 */
export class Prompt extends Message<Prompt> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string description = 3;
   */
  description = "";

  /**
   * @generated from field: string template = 4;
   */
  template = "";

  /**
   * @generated from field: google.protobuf.Struct metadata = 5;
   */
  metadata?: Struct;

  constructor(data?: PartialMessage<Prompt>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "prompt.v1.Prompt";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "template", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "metadata", kind: "message", T: Struct },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Prompt {
    return new Prompt().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Prompt {
    return new Prompt().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Prompt {
    return new Prompt().fromJsonString(jsonString, options);
  }

  static equals(a: Prompt | PlainMessage<Prompt> | undefined, b: Prompt | PlainMessage<Prompt> | undefined): boolean {
    return proto3.util.equals(Prompt, a, b);
  }
}

