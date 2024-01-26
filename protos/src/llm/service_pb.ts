// @generated by protoc-gen-es v1.6.0 with parameter "target=ts,import_extension="
// @generated from file llm/service.proto (package llm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from enum llm.v1.FinishReason
 */
export enum FinishReason {
  /**
   * @generated from enum value: NULL = 0;
   */
  NULL = 0,

  /**
   * @generated from enum value: LENGTH = 1;
   */
  LENGTH = 1,

  /**
   * @generated from enum value: STOP = 2;
   */
  STOP = 2,

  /**
   * @generated from enum value: ERROR = 3;
   */
  ERROR = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(FinishReason)
proto3.util.setEnumType(FinishReason, "llm.v1.FinishReason", [
  { no: 0, name: "NULL" },
  { no: 1, name: "LENGTH" },
  { no: 2, name: "STOP" },
  { no: 3, name: "ERROR" },
]);

/**
 * @generated from message llm.v1.Role
 */
export class Role extends Message<Role> {
  /**
   * @generated from oneof llm.v1.Role.role
   */
  role: {
    /**
     * @generated from field: string system = 1;
     */
    value: string;
    case: "system";
  } | {
    /**
     * @generated from field: string user = 2;
     */
    value: string;
    case: "user";
  } | {
    /**
     * @generated from field: string assistant = 3;
     */
    value: string;
    case: "assistant";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Role>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "llm.v1.Role";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "system", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "role" },
    { no: 2, name: "user", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "role" },
    { no: 3, name: "assistant", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "role" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Role {
    return new Role().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Role {
    return new Role().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Role {
    return new Role().fromJsonString(jsonString, options);
  }

  static equals(a: Role | PlainMessage<Role> | undefined, b: Role | PlainMessage<Role> | undefined): boolean {
    return proto3.util.equals(Role, a, b);
  }
}

/**
 * @generated from message llm.v1.ResponseFormat
 */
export class ResponseFormat extends Message<ResponseFormat> {
  /**
   * @generated from field: string type = 1;
   */
  type = "";

  constructor(data?: PartialMessage<ResponseFormat>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "llm.v1.ResponseFormat";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResponseFormat {
    return new ResponseFormat().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResponseFormat {
    return new ResponseFormat().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResponseFormat {
    return new ResponseFormat().fromJsonString(jsonString, options);
  }

  static equals(a: ResponseFormat | PlainMessage<ResponseFormat> | undefined, b: ResponseFormat | PlainMessage<ResponseFormat> | undefined): boolean {
    return proto3.util.equals(ResponseFormat, a, b);
  }
}

/**
 * @generated from message llm.v1.ChatMessage
 */
export class ChatMessage extends Message<ChatMessage> {
  /**
   * role of the message author. One of "system", "user", "assistant".
   *
   * @generated from field: string role = 1;
   */
  role = "";

  /**
   * content of the message
   *
   * @generated from field: string content = 2;
   */
  content = "";

  constructor(data?: PartialMessage<ChatMessage>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "llm.v1.ChatMessage";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "role", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "content", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ChatMessage {
    return new ChatMessage().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ChatMessage {
    return new ChatMessage().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ChatMessage {
    return new ChatMessage().fromJsonString(jsonString, options);
  }

  static equals(a: ChatMessage | PlainMessage<ChatMessage> | undefined, b: ChatMessage | PlainMessage<ChatMessage> | undefined): boolean {
    return proto3.util.equals(ChatMessage, a, b);
  }
}

/**
 * @generated from message llm.v1.TextCompletionParameters
 */
export class TextCompletionParameters extends Message<TextCompletionParameters> {
  /**
   * temperature of the sampling, between [0, 2]. default = 1.0
   *
   * @generated from field: optional float temperature = 1;
   */
  temperature?: number;

  /**
   * whether to stream partial completions back as they are generated. default = false
   *
   * @generated from field: optional bool stream = 2;
   */
  stream?: boolean;

  /**
   * @generated from field: optional uint32 top_k = 3;
   */
  topK?: number;

  /**
   * @generated from field: optional float top_p = 4;
   */
  topP?: number;

  /**
   * number of chat completion choices to generate for each input message. default = 1
   *
   * @generated from field: optional uint32 n = 5;
   */
  n?: number;

  /**
   * @generated from field: repeated string stop = 6;
   */
  stop: string[] = [];

  /**
   * @generated from field: optional uint32 max_tokens = 7;
   */
  maxTokens?: number;

  /**
   * @generated from field: optional float presence_penalty = 8;
   */
  presencePenalty?: number;

  /**
   * @generated from field: optional float frequency_penalty = 9;
   */
  frequencyPenalty?: number;

  /**
   * @generated from field: optional llm.v1.ResponseFormat response_format = 10;
   */
  responseFormat?: ResponseFormat;

  /**
   * @generated from field: optional uint32 seed = 11;
   */
  seed?: number;

  constructor(data?: PartialMessage<TextCompletionParameters>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "llm.v1.TextCompletionParameters";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "temperature", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 2, name: "stream", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 3, name: "top_k", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 4, name: "top_p", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 5, name: "n", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 6, name: "stop", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 7, name: "max_tokens", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
    { no: 8, name: "presence_penalty", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 9, name: "frequency_penalty", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 10, name: "response_format", kind: "message", T: ResponseFormat, opt: true },
    { no: 11, name: "seed", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TextCompletionParameters {
    return new TextCompletionParameters().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TextCompletionParameters {
    return new TextCompletionParameters().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TextCompletionParameters {
    return new TextCompletionParameters().fromJsonString(jsonString, options);
  }

  static equals(a: TextCompletionParameters | PlainMessage<TextCompletionParameters> | undefined, b: TextCompletionParameters | PlainMessage<TextCompletionParameters> | undefined): boolean {
    return proto3.util.equals(TextCompletionParameters, a, b);
  }
}

/**
 * @generated from message llm.v1.CompletionRequest
 */
export class CompletionRequest extends Message<CompletionRequest> {
  /**
   * @generated from field: string model = 1;
   */
  model = "";

  /**
   * a list of messages comprising all the conversation so far
   *
   * @generated from field: repeated llm.v1.ChatMessage messages = 2;
   */
  messages: ChatMessage[] = [];

  /**
   * @generated from field: optional llm.v1.TextCompletionParameters parameters = 3;
   */
  parameters?: TextCompletionParameters;

  constructor(data?: PartialMessage<CompletionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "llm.v1.CompletionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "model", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "messages", kind: "message", T: ChatMessage, repeated: true },
    { no: 3, name: "parameters", kind: "message", T: TextCompletionParameters, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CompletionRequest {
    return new CompletionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CompletionRequest {
    return new CompletionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CompletionRequest {
    return new CompletionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CompletionRequest | PlainMessage<CompletionRequest> | undefined, b: CompletionRequest | PlainMessage<CompletionRequest> | undefined): boolean {
    return proto3.util.equals(CompletionRequest, a, b);
  }
}

/**
 * @generated from message llm.v1.CompletionChoice
 */
export class CompletionChoice extends Message<CompletionChoice> {
  /**
   * index of the choice in the list of choices.
   *
   * @generated from field: uint32 index = 1;
   */
  index = 0;

  /**
   * message generated by the model.
   *
   * @generated from field: llm.v1.ChatMessage message = 2;
   */
  message?: ChatMessage;

  /**
   * @generated from field: string finish_reason = 3;
   */
  finishReason = "";

  constructor(data?: PartialMessage<CompletionChoice>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "llm.v1.CompletionChoice";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "index", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 2, name: "message", kind: "message", T: ChatMessage },
    { no: 3, name: "finish_reason", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CompletionChoice {
    return new CompletionChoice().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CompletionChoice {
    return new CompletionChoice().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CompletionChoice {
    return new CompletionChoice().fromJsonString(jsonString, options);
  }

  static equals(a: CompletionChoice | PlainMessage<CompletionChoice> | undefined, b: CompletionChoice | PlainMessage<CompletionChoice> | undefined): boolean {
    return proto3.util.equals(CompletionChoice, a, b);
  }
}

/**
 * @generated from message llm.v1.Usage
 */
export class Usage extends Message<Usage> {
  /**
   * number of tokens in the prompt.
   *
   * @generated from field: optional int32 prompt_tokens = 1;
   */
  promptTokens?: number;

  /**
   * number of tokens in the generated completion.
   *
   * @generated from field: optional int32 completion_tokens = 2;
   */
  completionTokens?: number;

  /**
   * total number of tokens used in the request (prompt + completion).
   *
   * @generated from field: optional int32 total_tokens = 3;
   */
  totalTokens?: number;

  constructor(data?: PartialMessage<Usage>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "llm.v1.Usage";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "prompt_tokens", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 2, name: "completion_tokens", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 3, name: "total_tokens", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Usage {
    return new Usage().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Usage {
    return new Usage().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Usage {
    return new Usage().fromJsonString(jsonString, options);
  }

  static equals(a: Usage | PlainMessage<Usage> | undefined, b: Usage | PlainMessage<Usage> | undefined): boolean {
    return proto3.util.equals(Usage, a, b);
  }
}

/**
 * @generated from message llm.v1.CompletionResponse
 */
export class CompletionResponse extends Message<CompletionResponse> {
  /**
   * unique id for the chat completion.
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * object type, which is always "chat.completion[.chunk]".
   *
   * @generated from field: string object = 2;
   */
  object = "";

  /**
   * unix timestamp (in seconds) of when the chat completion was created.
   *
   * @generated from field: uint64 created = 3;
   */
  created = protoInt64.zero;

  /**
   * model used for the completion
   *
   * @generated from field: string model = 4;
   */
  model = "";

  /**
   * list of generated completion choices for the input prompt
   *
   * @generated from field: repeated llm.v1.CompletionChoice choices = 5;
   */
  choices: CompletionChoice[] = [];

  /**
   * usage statistics for the completion request.
   *
   * @generated from field: llm.v1.Usage usage = 6;
   */
  usage?: Usage;

  /**
   * @generated from field: string system_fingerprint = 7;
   */
  systemFingerprint = "";

  constructor(data?: PartialMessage<CompletionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "llm.v1.CompletionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "object", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "created", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 4, name: "model", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "choices", kind: "message", T: CompletionChoice, repeated: true },
    { no: 6, name: "usage", kind: "message", T: Usage },
    { no: 7, name: "system_fingerprint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CompletionResponse {
    return new CompletionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CompletionResponse {
    return new CompletionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CompletionResponse {
    return new CompletionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CompletionResponse | PlainMessage<CompletionResponse> | undefined, b: CompletionResponse | PlainMessage<CompletionResponse> | undefined): boolean {
    return proto3.util.equals(CompletionResponse, a, b);
  }
}

