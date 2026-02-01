const std = @import("std");
const solution = @import("solution.zig");

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    const input = try std.fs.cwd().readFileAlloc(
        allocator,
        "../input.txt",
        10 * 1024 * 1024,
    );
    defer allocator.free(input);

    const result = try solution.solve(input, allocator);

    const stdout = std.io.getStdOut().writer();
    try stdout.print("{any}\n", .{result});
}

