const std = @import("std");
const aoclib = @import("aoclib");

pub fn solve(input: []const u8, allocator: std.mem.Allocator) !usize {
    _ = allocator;

    // Example: count non-empty lines
    var count: usize = 0;
    var it = std.mem.splitScalar(u8, input, '\n');
    while (it.next()) |line| {
        if (line.len > 0) count += 1;
    }
    return count;
}

test "basic example" {
    const input =
        \\a
        \\b
        \\c
    ;
    const result = try solve(input, std.testing.allocator);
    try std.testing.expectEqual(@as(usize, 3), result);
}

