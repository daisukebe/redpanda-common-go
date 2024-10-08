// @generated by protoc-gen-es v1.6.0 with parameter "target=ts,import_extension=,js_import_style=legacy_commonjs"
// @generated from file redpanda/api/common/v1alpha1/pagination.proto (package redpanda.api.common.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * KeySetPageToken represents a pagination token for KeySet pagination.
 * It marks the beginning of a page where records start from the key that
 * satisfies the condition key >= value_greater_equal. Records are sorted
 * alphabetically by key in ascending order.
 *
 * @generated from message redpanda.api.common.v1alpha1.KeySetPageToken
 */
export class KeySetPageToken extends Message<KeySetPageToken> {
  /**
   * @generated from field: string key = 1;
   */
  key = "";

  /**
   * @generated from field: string value_greater_equal = 2;
   */
  valueGreaterEqual = "";

  constructor(data?: PartialMessage<KeySetPageToken>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.common.v1alpha1.KeySetPageToken";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "value_greater_equal", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): KeySetPageToken {
    return new KeySetPageToken().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): KeySetPageToken {
    return new KeySetPageToken().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): KeySetPageToken {
    return new KeySetPageToken().fromJsonString(jsonString, options);
  }

  static equals(a: KeySetPageToken | PlainMessage<KeySetPageToken> | undefined, b: KeySetPageToken | PlainMessage<KeySetPageToken> | undefined): boolean {
    return proto3.util.equals(KeySetPageToken, a, b);
  }
}

