const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    // Shared AoC utils
    const aoclib = b.addModule("aoclib", .{
        .source_file = b.path("../../../libs/zig/aoclib/src/aoclib.zig"),
    });

    // Main executable
    const exe = b.addExecutable(.{
        .name = "day",
        .root_source_file = b.path("src/main.zig"),
        .target = target,
        .optimize = optimize,
    });
    exe.root_module.addImport("aoclib", aoclib);
    b.installArtifact(exe);

    // --- run ---
    const run_cmd = b.addRunArtifact(exe);
    if (b.args) |args| run_cmd.addArgs(args);

    const run_step = b.step("run", "Run solution");
    run_step.dependOn(&run_cmd.step);

    // --- tests ---
    const tests = b.addTest(.{
        .root_source_file = b.path("src/solution.zig"),
        .target = target,
        .optimize = optimize,
    });
    tests.root_module.addImport("aoclib", aoclib);

    const test_step = b.step("test", "Run tests");
    test_step.dependOn(&tests.step);

    // --- benchmark ---
    const bench_step = b.step("benchmark", "Run benchmarks");
    bench_step.dependOn(&exe.step);
}

