const std = @import("std");
const print = std.debug.print;

pub fn main() !void {
    const test_creditcard_num = "4532591760818986";
    if (verify(test_creditcard_num)) {
        print("Valid :3\n", .{});
    }
    // very nice
    else {
        print("Invalid >:(\n", .{});
    }
}

pub fn verify(cc: *const [16:0]u8) bool {

    // the entire algorithm is about if this sum is divisible by 10
    // if yes your card is valid, if no... jail for fraud :3
    var final_sum: u8 = 0;

    const input_string: *const [16:0]u8 = cc;
    const input_split = std.mem.splitAny(u8, input_string, "");

    for (0.., input_split.buffer) |i, value| {
        // just a test print
        // print("index[{}]: {}\n", .{ i, value - 48 });

        // we dont use any fancy library to turn "1" into 1
        // an ASCII trick does great, ascii value of "1" is 49 so we just take away 48
        const digit = value - 48;

        // handle even indexes
        if (i % 2 == 0) {

            // example digit is 6
            // print("\n{} EVEN", .{digit});

            // 6 * 2 -> 12
            const number_doubled: f16 = @floatFromInt(digit * 2);
            
            // debug print
            // print("\nThe number doubled is {} \n", .{number_doubled});

            // 12 -> 2
            const first_digit = @mod(number_doubled, 10);
            // 12 -> 1.2 -> 1
            const second_digit = @floor(number_doubled / 10);
            // 1 + 2
            const digit_sum = first_digit + second_digit;
            // 3
            // print("the sum of the digits is: {}", .{digit_sum});

            // we do a little type casting float -> int
            final_sum += @intFromFloat(digit_sum);
        }

        // handle odd indexes
        else {
            // print("\n{} ODD", .{digit});
            final_sum += digit;
        }
    }
    // print("final sum is: {} \n", .{final_sum});

    // beauty of zig
    if (@mod(final_sum, 10) == 0) return true else return false;
}
